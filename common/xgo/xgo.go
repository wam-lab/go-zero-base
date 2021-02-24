package xgo

import "github.com/tal-tech/go-zero/core/logx"

func Go(runner func())  {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				logx.Error(err)
			}
		}()

		runner()
	}()
}