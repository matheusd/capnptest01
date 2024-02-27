package capnptest01

import (
	"testing"

	"capnproto.org/go/capnp/v3"
)

func BenchmarkSetText01(b *testing.B) {
	arena := capnp.SingleSegment(nil)
	_, seg, err := capnp.NewMessage(arena)
	if err != nil {
		b.Fatal(err)
	}

	tx, err := NewTransaction(seg)
	if err != nil {
		b.Fatal(err)
	}

	err = tx.SetDescription("my own descr")
	if err != nil {
		b.Fatal(err)
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		err = tx.SetDescription("my own descr")
		if err != nil {
			b.Fatal(err)
		}
	}

	// b.Log(arena.String())
}

func BenchmarkSetText02(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		arena := capnp.SingleSegment(nil)
		_, seg, err := capnp.NewMessage(arena)
		if err != nil {
			b.Fatal(err)
		}

		tx, err := NewRootTransaction(seg)
		if err != nil {
			b.Fatal(err)
		}

		err = tx.SetDescription("my own descr")
		if err != nil {
			b.Fatal(err)
		}

		arena.Release()
	}
}

func BenchmarkSetText03(b *testing.B) {
	var msg capnp.Message
	arena := capnp.SingleSegment(nil)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		seg, err := msg.Reset(arena)
		if err != nil {
			b.Fatal(err)
		}

		tx, err := NewTransaction(seg)
		if err != nil {
			b.Fatal(err)
		}

		err = tx.SetDescription("my own descr")
		if err != nil {
			b.Fatal(err)
		}
	}

	// b.Log(arena.String())
}

func BenchmarkSetText04(b *testing.B) {
	buf := make([]byte, 0, 2048)
	arena := capnp.SingleSegment(buf)

	b.ReportAllocs()
	b.ResetTimer()

	msg := capnp.Message{
		Arena: arena,
	}

	for i := 0; i < b.N; i++ {
		seg, err := msg.Segment(0)
		if err != nil {
			b.Fatal(err)
		}

		tx, err := NewRootTransaction(seg)
		if err != nil {
			b.Fatal(err)
		}

		err = tx.SetDescription("my own descr")
		if err != nil {
			b.Fatal(err)
		}
	}

	// b.Log(arena.String())
}
