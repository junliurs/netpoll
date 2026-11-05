package netpoll

import (
	"bytes"
	"testing"
)

func TestMeshlibBorrowConsumedMessagePreservesNodesAndBoundary(t *testing.T) {
	b := NewLinkBuffer()
	defer b.Close()
	first := []byte("first-message")
	second := []byte("next-message")
	_, _ = b.WriteBinary(first)
	_ = b.Flush()
	mark := b.MarkRead()
	if err := b.Skip(len(first)); err != nil {
		t.Fatal(err)
	}
	_, _ = b.WriteBinary(second)
	_ = b.Flush()
	var scratch [4][]byte
	parts := b.ConsumedSlices(mark, scratch[:0])
	if got := b.ConsumedLen(mark); got != len(first) {
		t.Fatalf("consumed %d bytes, want %d", got, len(first))
	}
	if len(parts) != 1 || !bytes.Equal(parts[0], first) {
		t.Fatal(parts)
	}
	firstView := parts[0]
	again := b.ConsumedSlices(mark, scratch[:0])
	if &again[0][0] != &firstView[0] {
		t.Fatal("copied payload")
	}
	remaining, err := b.Next(len(second))
	if err != nil || !bytes.Equal(remaining, second) {
		t.Fatal("consumed next message", err)
	}
	if got := b.ConsumedLen(mark); got != len(first)+len(second) {
		t.Fatalf("consumed %d bytes across nodes", got)
	}
}
