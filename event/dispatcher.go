package event

type Dispatcher[T any] struct {
	callbacks []func(e T) bool
	handled   bool
}

func (d *Dispatcher[T]) AddCallBack(cb func(e T) bool) {
	d.callbacks = append(d.callbacks, cb)
}

func (d *Dispatcher[T]) Dispatch(e T) {
	d.handled = false
	for _, cb := range d.callbacks {
		if !d.handled {
			d.handled = cb(e)
		}
	}
}
