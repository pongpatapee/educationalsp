package lsp

type TextDocumentItem struct {
	/**
	 * The text document's URI.
	 */
	URI string `json:"uri"`

	/**
	 * The text document's language identifier.
	 */
	LanguageID string `json:"languageId"`

	/**
	 * The version number of this document (it will increase after each
	 * change, including undo/redo).
	 */
	Version int `json:"version"`

	/**
	 * The content of the opened text document.
	 */
	Text string `json:"text"`
}

type TextDocumentIdentifer struct {
	URI string `json:"uri"`
}

type VersionTextDocumentIdentifer struct {
	TextDocumentIdentifer
	version int `json:"version"`
}

type TextDocumentPositionParams struct {
	TextDocument TextDocumentIdentifer `json:"textDocument"`
	Position     Position              `json:"position"`
}

type Position struct {
	Line      int `json:"line"`
	Character int `json:"character"`
}
