package protocol

type StoreNodeRequestConfig struct {
	WaitForResponse   bool
	StopWhenDataFound bool
	InitialPageSize   uint32
	FurtherPageSize   uint32
}

type StoreNodeRequestOption func(*StoreNodeRequestConfig)

func defaultStoreNodeRequestConfig() StoreNodeRequestConfig {
	return StoreNodeRequestConfig{
		WaitForResponse:   true,
		StopWhenDataFound: true,
		InitialPageSize:   initialStoreNodeRequestPageSize,
		FurtherPageSize:   defaultStoreNodeRequestPageSize,
	}
}

func buildStoreNodeRequestConfig(opts []StoreNodeRequestOption) StoreNodeRequestConfig {
	cfg := defaultStoreNodeRequestConfig()

	// TODO: remove these 2 when fixed: https://github.com/waku-org/nwaku/issues/2317
	opts = append(opts, WithStopWhenDataFound(false))
	opts = append(opts, WithInitialPageSize(defaultStoreNodeRequestPageSize))

	for _, opt := range opts {
		opt(&cfg)
	}

	return cfg
}

func WithWaitForResponseOption(waitForResponse bool) StoreNodeRequestOption {
	return func(c *StoreNodeRequestConfig) {
		c.WaitForResponse = waitForResponse
	}
}

func WithStopWhenDataFound(stopWhenDataFound bool) StoreNodeRequestOption {
	return func(c *StoreNodeRequestConfig) {
		c.StopWhenDataFound = stopWhenDataFound
	}
}

func WithInitialPageSize(initialPageSize uint32) StoreNodeRequestOption {
	return func(c *StoreNodeRequestConfig) {
		c.InitialPageSize = initialPageSize
	}
}

func WithFurtherPageSize(furtherPageSize uint32) StoreNodeRequestOption {
	return func(c *StoreNodeRequestConfig) {
		c.FurtherPageSize = furtherPageSize
	}
}