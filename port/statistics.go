package port

type Statics interface {
	/**
	invoke peer secs
	 */
	StaticInfo(copyBytes int64, port *port)
}
