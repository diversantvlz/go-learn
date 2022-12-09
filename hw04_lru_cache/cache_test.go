package hw04lrucache

import (
	"math/rand"
	"strconv"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCache(t *testing.T) {
	t.Run("empty cache", func(t *testing.T) {
		c := NewCache(10)

		_, ok := c.Get("aaa")
		require.False(t, ok)

		_, ok = c.Get("bbb")
		require.False(t, ok)
	})

	t.Run("simple", func(t *testing.T) {
		c := NewCache(5)

		wasInCache := c.Set("aaa", 100)
		require.False(t, wasInCache)

		wasInCache = c.Set("bbb", 200)
		require.False(t, wasInCache)

		val, ok := c.Get("aaa")
		require.True(t, ok)
		require.Equal(t, 100, val)

		val, ok = c.Get("bbb")
		require.True(t, ok)
		require.Equal(t, 200, val)

		wasInCache = c.Set("aaa", 300)
		require.True(t, wasInCache)

		val, ok = c.Get("aaa")
		require.True(t, ok)
		require.Equal(t, 300, val)

		val, ok = c.Get("ccc")
		require.False(t, ok)
		require.Nil(t, val)
	})

	t.Run("purge logic", func(t *testing.T) {
		c := NewCache(3)
		dataset := map[string]int{
			"aaa": 100,
			"bbb": 200,
			"ccc": 300,
		}

		for key, value := range dataset {
			c.Set(Key(key), value)
			val, ok := c.Get(Key(key))
			require.True(t, ok)
			require.Equal(t, value, val)
		}

		c.Clear()

		for key := range dataset {
			_, ok := c.Get(Key(key))
			require.False(t, ok)
		}
	})

	t.Run("oversize logic", func(t *testing.T) {
		c := NewCache(2)
		c.Set("a", 1)
		c.Set("b", 2)
		c.Set("c", 3)

		_, notOk := c.Get("aaa")
		require.False(t, notOk)
	})

	t.Run("oldest logic", func(t *testing.T) {
		c := NewCache(3)
		c.Set("a", 1)
		c.Set("b", 2)
		c.Set("c", 3)
		c.Get("a")
		c.Get("c")
		c.Set("d", 4)

		_, notOk := c.Get("b")
		require.False(t, notOk)
	})
}

func TestCacheMultithreading(t *testing.T) {
	c := NewCache(10)
	wg := &sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < 1_000_000; i++ {
			c.Set(Key(strconv.Itoa(i)), i)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 1_000_000; i++ {
			c.Get(Key(strconv.Itoa(rand.Intn(1_000_000))))
		}
	}()

	wg.Wait()
}
