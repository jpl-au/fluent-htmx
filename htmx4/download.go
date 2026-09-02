package htmx

import "net/http"

// Download extension, bundled in htmax.js (with the core htmx.js build, include the download
// extension script yourself). It streams a response to the browser as a file, firing
// htmx:download:start, htmx:download:progress and htmx:download:complete events as it
// goes. A download is triggered three ways: the swap.Download swap style on the requesting
// element and a Content-Disposition: attachment response both stream the response itself
// and skip the swap; the HX-Download response header set by HxDownload fetches a separate
// URL while the response goes through the swap as usual.

// HxDownload tells the client to fetch the given URL and save it through the browser's
// download mechanism, by setting the HX-Download response header. The extension starts
// that fetch, and the response itself is swapped as usual, so the body is the message to
// show while the download runs - "Your download has started", say. A response with no
// body clears the target under the default swap; send HxReswap(w, swap.None) with it to
// leave the page as it is.
func HxDownload(w http.ResponseWriter, url string) {
	w.Header().Set(HXDownloadHeader, url)
}
