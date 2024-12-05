package redis

type Redis struct{}

func New() (*Redis, error) {
	return &Redis{}, nil
}

func (r *Redis) Close() {
}
