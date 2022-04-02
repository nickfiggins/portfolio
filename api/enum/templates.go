//go:generate enumer -type=Template -json -transform=kebab
package enum

type Template int

const(
	TutoringFormConfirm Template = iota
	GeneralFormConfirm
	FreelanceFormConfirm
)
