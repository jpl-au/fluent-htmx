package htmx

import "net/http"

// Download extension, bundled in htmax.js (with the core htmx.js build, include the download
// extension script yourself). It streams a response to the browser as
// a file instead of swapping it into the DOM, firing htmx:download:start, htmx:download:progress
// and htmx:download:complete events as it goes. A download is triggered three ways: the
// swap.Download swap style on the requesting element, a Content-Disposition: attachment
// response, or the HX-Download response header set by HxDownload.

// HxDownload tells the client to download the given URL as a file rather than swapping the
// response into the page. It sets the HX-Download response header; the bundled download
// extension then fetches that URL and saves it through the browser's download mechanism.
func HxDownload(w http.ResponseWriter, url string) {
	w.Header().Set(HXDownloadHeader, url)
}
