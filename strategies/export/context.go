package export

type ExportContext struct {
	exportInterface exportInterface
}

func InitExportContext(e exportInterface) *ExportContext {
	return &ExportContext {
		exportInterface: e,
	}
}

func (c *ExportContext) SetExportInterface(e exportInterface) {
	c.exportInterface = e
}

func (c *ExportContext) Export() {
	c.exportInterface.export()
}