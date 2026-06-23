package chain

import (
	"context"
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

// maxBlockSpan caps the eth_getLogs range scanned per filter call. Many RPC
// providers reject getLogs over wide ranges, so a long catch-up (e.g. after
// RPC downtime) is chunked into provider-friendly windows.
const maxBlockSpan = 2000

// WatchRecordEvents polls the contract for Registered / Transferred /
// RecordSet / RecordRemoved events and delivers them to handler. Polling
// (rather than a WS subscription) works against any RPC transport, including
// plain-HTTP Hardhat (HLD open issue 5: TTL + push invalidation). Blocks until
// ctx is cancelled.
//
// last only ever advances past a block range that was filtered successfully
// for every event type: if any FilterXxx call fails (a transient RPC error),
// the range is left to be retried on the next tick rather than silently
// dropping the invalidation events it contained.
func (c *Client) WatchRecordEvents(ctx context.Context, interval time.Duration, handler func(RecordEvent)) error {
	last, err := c.ChainHead(ctx)
	if err != nil {
		return err
	}
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-ticker.C:
		}

		head, err := c.ChainHead(ctx)
		if err != nil {
			if ctx.Err() != nil {
				return ctx.Err()
			}
			c.log.Warn("event poll: head", "err", err)
			continue
		}
		for last < head {
			end := head
			if end-last > maxBlockSpan {
				end = last + maxBlockSpan
			}
			opts := &bind.FilterOpts{Start: last + 1, End: &end, Context: ctx}
			if err := c.collectEvents(opts, handler); err != nil {
				// Leave last unchanged so [last+1, end] is retried next tick.
				c.log.Warn("event poll: collect", "from", last+1, "to", end, "err", err)
				break
			}
			last = end
		}
	}
}

// collectEvents filters all four event types over opts and delivers them to
// handler. It returns the first error encountered (filter call or iterator
// error) so the caller can avoid advancing past an incompletely-scanned range.
func (c *Client) collectEvents(opts *bind.FilterOpts, handler func(RecordEvent)) error {
	itSet, err := c.dapp.FilterRecordSet(opts, nil, nil)
	if err != nil {
		return fmt.Errorf("filter RecordSet: %w", err)
	}
	for itSet.Next() {
		handler(RecordEvent{
			Kind: EventRecordSet, Name: itSet.Event.Name, NameHash: itSet.Event.NameHash,
			RecordType: itSet.Event.RecordType, Selector: itSet.Event.Selector,
			Block: itSet.Event.Raw.BlockNumber,
		})
	}
	if err := itSet.Error(); err != nil {
		itSet.Close()
		return fmt.Errorf("iterate RecordSet: %w", err)
	}
	itSet.Close()

	itRem, err := c.dapp.FilterRecordRemoved(opts, nil, nil)
	if err != nil {
		return fmt.Errorf("filter RecordRemoved: %w", err)
	}
	for itRem.Next() {
		handler(RecordEvent{
			Kind: EventRecordRemoved, Name: itRem.Event.Name, NameHash: itRem.Event.NameHash,
			RecordType: itRem.Event.RecordType, Selector: itRem.Event.Selector,
			Block: itRem.Event.Raw.BlockNumber,
		})
	}
	if err := itRem.Error(); err != nil {
		itRem.Close()
		return fmt.Errorf("iterate RecordRemoved: %w", err)
	}
	itRem.Close()

	itXfer, err := c.dapp.FilterTransferred(opts, nil, nil, nil)
	if err != nil {
		return fmt.Errorf("filter Transferred: %w", err)
	}
	for itXfer.Next() {
		handler(RecordEvent{
			Kind: EventTransferred, NameHash: itXfer.Event.NameHash,
			Block: itXfer.Event.Raw.BlockNumber,
		})
	}
	if err := itXfer.Error(); err != nil {
		itXfer.Close()
		return fmt.Errorf("iterate Transferred: %w", err)
	}
	itXfer.Close()

	itReg, err := c.dapp.FilterRegistered(opts, nil, nil)
	if err != nil {
		return fmt.Errorf("filter Registered: %w", err)
	}
	for itReg.Next() {
		handler(RecordEvent{
			Kind: EventRegistered, Name: itReg.Event.Name, NameHash: itReg.Event.NameHash,
			Block: itReg.Event.Raw.BlockNumber,
		})
	}
	if err := itReg.Error(); err != nil {
		itReg.Close()
		return fmt.Errorf("iterate Registered: %w", err)
	}
	itReg.Close()

	return nil
}
