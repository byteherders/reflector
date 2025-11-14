package server

import "html/template"

var reflectionTemplate = template.Must(template.New("page").Parse(pageTemplateHTML))

const pageTemplateHTML = `<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="utf-8">
	<meta name="viewport" content="width=device-width, initial-scale=1">
	<title>HTTP Reflector</title>
	<link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bootstrap@5.3.2/dist/css/bootstrap.min.css">
	<style>
		body { background-color: #f8f9fa; }
		pre { background-color: #0f172a; color: #e2e8f0; padding: 1rem; border-radius: 0.5rem; }
		code { font-size: 0.875rem; }
	</style>
</head>
<body>
	<div class="container py-4">
		<header class="mb-4">
			<div class="d-flex flex-wrap justify-content-between align-items-center gap-3">
				<div>
					<h1 class="h3 mb-1">HTTP Reflector</h1>
					<p class="text-muted mb-0">Observing request from <code>{{.Reflection.RemoteAddr}}</code></p>
				</div>
				<span class="badge text-bg-secondary">{{.Reflection.Timestamp}}</span>
			</div>
		</header>

		<div id="status-message" class="alert alert-{{.StatusVariant}} mb-4" role="alert">
			{{.StatusMessage}}
		</div>

		<section class="mb-4">
			<div class="card shadow-sm">
				<div class="card-header fw-semibold">Request Overview</div>
				<div class="card-body">
					<div class="row gy-3">
						<div class="col-md-6">
							<dl class="row mb-0 small">
								<dt class="col-sm-4 text-muted">Method</dt>
								<dd class="col-sm-8">{{.Reflection.Method}}</dd>
								<dt class="col-sm-4 text-muted">Protocol</dt>
								<dd class="col-sm-8">{{.Reflection.Proto}}</dd>
								<dt class="col-sm-4 text-muted">Scheme</dt>
								<dd class="col-sm-8">{{.Reflection.Scheme}}</dd>
								<dt class="col-sm-4 text-muted">Host</dt>
								<dd class="col-sm-8">{{.Reflection.Host}}</dd>
								<dt class="col-sm-4 text-muted">Request URI</dt>
								<dd class="col-sm-8"><code>{{.Reflection.RequestURI}}</code></dd>
							</dl>
						</div>
						<div class="col-md-6">
							<dl class="row mb-0 small">
								<dt class="col-sm-5 text-muted">Remote Address</dt>
								<dd class="col-sm-7"><code>{{.Reflection.RemoteAddr}}</code></dd>
								<dt class="col-sm-5 text-muted">Remote IP</dt>
								<dd class="col-sm-7"><code>{{.Reflection.RemoteIP}}</code></dd>
								<dt class="col-sm-5 text-muted">Remote Port</dt>
								<dd class="col-sm-7">{{.Reflection.RemotePort}}</dd>
								<dt class="col-sm-5 text-muted">Content Length</dt>
								<dd class="col-sm-7">{{.Reflection.ContentLength}}</dd>
								<dt class="col-sm-5 text-muted">Transfer Encoding</dt>
								<dd class="col-sm-7">
									{{if .Reflection.TransferEncoding}}
										{{range .Reflection.TransferEncoding}}<span class="badge text-bg-secondary me-1">{{.}}</span>{{end}}
									{{else}}
										<span class="text-muted">none</span>
									{{end}}
								</dd>
							</dl>
						</div>
					</div>
				</div>
			</div>
		</section>

		<section class="mb-4">
			<div class="row g-4">
				<div class="col-lg-6">
					<div class="card h-100 shadow-sm">
						<div class="card-header fw-semibold">Headers</div>
						<div class="card-body">
							{{if .Headers}}
								<div class="table-responsive">
									<table class="table table-sm align-middle mb-0">
										<tbody>
											{{range .Headers}}
											<tr>
												<th scope="row" class="text-nowrap">{{.Key}}</th>
												<td>
													{{range .Values}}<span class="badge text-bg-primary me-1">{{.}}</span>{{end}}
												</td>
											</tr>
											{{end}}
										</tbody>
									</table>
								</div>
							{{else}}
								<p class="text-muted mb-0">No headers were supplied.</p>
							{{end}}
						</div>
					</div>
				</div>
				<div class="col-lg-6">
					<div class="card h-100 shadow-sm">
						<div class="card-header fw-semibold">Query Parameters</div>
						<div class="card-body">
							{{if .Query}}
								<div class="table-responsive">
									<table class="table table-sm align-middle mb-0">
										<tbody>
											{{range .Query}}
											<tr>
												<th scope="row" class="text-nowrap">{{.Key}}</th>
												<td>
													{{range .Values}}<span class="badge text-bg-info me-1">{{.}}</span>{{end}}
												</td>
											</tr>
											{{end}}
										</tbody>
									</table>
								</div>
							{{else}}
								<p class="text-muted mb-0">No query parameters detected.</p>
							{{end}}
						</div>
					</div>
				</div>
			</div>
		</section>

		<section class="mb-4">
			<div class="row g-4">
				<div class="col-lg-6">
					<div class="card h-100 shadow-sm">
						<div class="card-header fw-semibold">Cookies</div>
						<div class="card-body">
							{{if .Reflection.Cookies}}
								<div class="table-responsive">
									<table class="table table-sm align-middle mb-0">
										<tbody>
											{{range .Reflection.Cookies}}
											<tr>
												<th scope="row" class="text-nowrap">{{.Name}}</th>
												<td><code>{{.Value}}</code></td>
											</tr>
											{{end}}
										</tbody>
									</table>
								</div>
							{{else}}
								<p class="text-muted mb-0">No cookies were provided.</p>
							{{end}}
						</div>
					</div>
				</div>
				<div class="col-lg-6">
					<div class="card h-100 shadow-sm">
						<div class="card-header fw-semibold">TLS</div>
						<div class="card-body">
							{{with .Reflection.TLS}}
								<dl class="row mb-0 small">
									<dt class="col-sm-4 text-muted">Version</dt>
									<dd class="col-sm-8">{{.Version}}</dd>
									<dt class="col-sm-4 text-muted">Cipher Suite</dt>
									<dd class="col-sm-8">{{.CipherSuite}}</dd>
									<dt class="col-sm-4 text-muted">Server Name</dt>
									<dd class="col-sm-8">{{if .ServerName}}{{.ServerName}}{{else}}<span class="text-muted">n/a</span>{{end}}</dd>
									<dt class="col-sm-4 text-muted">ALPN</dt>
									<dd class="col-sm-8">{{if .Negotiated}}{{.Negotiated}}{{else}}<span class="text-muted">n/a</span>{{end}}</dd>
								</dl>
							{{else}}
								<p class="text-muted mb-0">Connection is not using TLS.</p>
							{{end}}
						</div>
					</div>
				</div>
			</div>
		</section>

		<section class="mb-4">
			<div class="card shadow-sm">
				<div class="card-header fw-semibold d-flex justify-content-between align-items-center">
					<span>Request Body</span>
					<span class="text-muted small">{{len .Reflection.BodyPreview}} bytes shown</span>
				</div>
				<div class="card-body">
					{{if .Reflection.BodyPreview}}
						<pre class="mb-0">{{.Reflection.BodyPreview}}</pre>
					{{else}}
						<p class="text-muted mb-0">No request body captured.</p>
					{{end}}
				</div>
			</div>
		</section>

		<section class="mb-5">
			<div class="card shadow-sm">
				<div class="card-header fw-semibold">Browser Metadata</div>
				<div class="card-body">
					{{if .ClientJSON}}
						<pre class="mb-0">{{.ClientJSON}}</pre>
					{{else}}
						<p class="text-muted mb-0">Waiting for the browser script to provide additional context...</p>
					{{end}}
				</div>
			</div>
		</section>

		<footer class="text-muted small">
			HTTP Reflector · helpful for CDN debugging and origin verification.
		</footer>
	</div>

	<script>
		window.__reflectorHasClientData = {{if .HasClientData}}true{{else}}false{{end}};
		{{.ClientScript}}
	</script>
</body>
</html>`

const clientCollectorScript = `(function () {
	"use strict";
	const hasClientData = !!window.__reflectorHasClientData;
	const statusEl = document.getElementById("status-message");
	const endpoint = "/collect";

	function updateStatus(message, variant) {
		if (!statusEl) {
			return;
		}
		statusEl.textContent = message;
		statusEl.className = "alert alert-" + variant;
	}

	if (!window.fetch) {
		updateStatus("Browser fetch API is unavailable in this environment.", "warning");
		return;
	}

	if (hasClientData) {
		return;
	}

	function storageKeys(storage) {
		if (!storage) {
			return [];
		}
		const keys = [];
		try {
			for (let i = 0; i < storage.length; i += 1) {
				const key = storage.key(i);
				if (key !== null) {
					keys.push(key);
				}
			}
		} catch (err) {
			return ["unavailable"];
		}
		return keys;
	}

	function boolOrNull(value) {
		return typeof value === "boolean" ? value : null;
	}

	function gatherClientData() {
		const nav = typeof navigator === "undefined" ? {} : navigator;
		const scr = typeof window === "undefined" || !window.screen ? {} : window.screen;
		const connection = nav.connection || nav.mozConnection || nav.webkitConnection;
		const colorScheme = typeof window !== "undefined" && window.matchMedia
			? (window.matchMedia("(prefers-color-scheme: dark)").matches
				? "dark"
				: (window.matchMedia("(prefers-color-scheme: light)").matches ? "light" : "no-preference"))
			: null;
		const touchPoints = typeof nav.maxTouchPoints === "number" ? nav.maxTouchPoints : null;
		let timezone = null;
		try {
			if (typeof Intl !== "undefined" && Intl.DateTimeFormat) {
				timezone = Intl.DateTimeFormat().resolvedOptions().timeZone || null;
			}
		} catch (_) {
			timezone = null;
		}

		const connectionInfo = connection
			? {
				type: connection.type || null,
				effectiveType: connection.effectiveType || null,
				downlink: connection.downlink || null,
				rtt: connection.rtt || null,
				saveData: boolOrNull(connection.saveData),
			}
			: null;

		const performanceInfo = typeof performance !== "undefined" && performance.memory
			? {
				jsHeapSizeLimit: performance.memory.jsHeapSizeLimit,
				totalJSHeapSize: performance.memory.totalJSHeapSize,
				usedJSHeapSize: performance.memory.usedJSHeapSize,
			}
			: null;

		return {
			timestamp: new Date().toISOString(),
			location: window.location.href,
			referrer: document.referrer || null,
			userAgent: nav.userAgent || null,
			language: nav.language || null,
			languages: Array.isArray(nav.languages) ? nav.languages : [],
			platform: nav.platform || null,
			deviceMemory: typeof nav.deviceMemory === "number" ? nav.deviceMemory : null,
			hardwareConcurrency: typeof nav.hardwareConcurrency === "number" ? nav.hardwareConcurrency : null,
			doNotTrack: nav.doNotTrack || window.doNotTrack || null,
			cookieEnabled: boolOrNull(nav.cookieEnabled),
			onLine: boolOrNull(nav.onLine),
			timezone: timezone,
			timezoneOffsetMinutes: new Date().getTimezoneOffset(),
			screen: {
				width: scr.width || null,
				height: scr.height || null,
				availWidth: scr.availWidth || null,
				availHeight: scr.availHeight || null,
				colorDepth: scr.colorDepth || null,
				pixelDepth: scr.pixelDepth || null,
			},
			viewport: {
				width: typeof window.innerWidth === "number" ? window.innerWidth : null,
				height: typeof window.innerHeight === "number" ? window.innerHeight : null,
			},
			colorScheme: colorScheme,
			historyLength: window.history ? window.history.length : null,
			connection: connectionInfo,
			performance: performanceInfo,
			localStorageKeys: storageKeys(window.localStorage),
			sessionStorageKeys: storageKeys(window.sessionStorage),
			touchSupport: {
				maxTouchPoints: touchPoints,
				touchEvent: typeof window.ontouchstart !== "undefined",
				pointerEvent: typeof window.PointerEvent !== "undefined",
			},
			mediaDevices: !!(nav.mediaDevices && nav.mediaDevices.enumerateDevices),
		};
	}

	function onReady(callback) {
		if (document.readyState === "loading") {
			document.addEventListener("DOMContentLoaded", callback, { once: true });
		} else {
			callback();
		}
	}

	async function send() {
		updateStatus("Collecting additional details from your browser...", "info");
		const payload = gatherClientData();
		try {
			const response = await fetch(endpoint, {
				method: "POST",
				credentials: "same-origin",
				headers: {
					"Content-Type": "application/json",
				},
				body: JSON.stringify(payload),
			});
			if (!response.ok) {
				const text = await response.text();
				throw new Error("server responded with " + response.status + ": " + text);
			}
			updateStatus("Rendering browser metadata...", "info");
			const html = await response.text();
			document.open();
			document.write(html);
			document.close();
		} catch (err) {
			const message = err && err.message ? err.message : String(err);
			updateStatus("Failed to capture browser details: " + message, "danger");
		}
	}

	onReady(send);
})()`
