package common

// "print_on_side" key is used to set the type of side printing for canvases. There are three possible values:
// "regular" - to extend print area to the sides of canvas
// "mirror" - to keep original print area and mirror it to the sides
// "off" - stop printing on sides
type PrintDetails struct {
	PrintOnSide PrintOnSideEnum `json:"print_on_side"`
}

type PrintOnSideEnum string

const (
	PrintOnSideRegular PrintOnSideEnum = "regular"
	PrintOnSideMirror  PrintOnSideEnum = "mirror"
	PrintOnSideOff     PrintOnSideEnum = "off"
)
