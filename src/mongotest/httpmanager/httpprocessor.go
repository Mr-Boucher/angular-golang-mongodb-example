package httpmanager

type Processor interface {
	Execute( HttpContext )
}
