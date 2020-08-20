/*
 * Domeneshop API Documentation
 *
 * # Overview Domeneshop offers a simple, REST-based API, which currently supports the following  features: - List domains - List invoices - Create, read, update and delete DNS records for domains - Create, read, update and delete HTTP forwards (\"WWW forwarding\") for domains - Dynamic DNS (DDNS) update endpoints for use in consumer routers  More features are planned, including: - Web hosting administration - Email address and email user/account administration  # Authentication The Domeneshop API currently supports only one method of authentication, **HTTP Basic Auth**. More authentication methods may be added in the future.  To generate credentials, visit <a href=\"https://www.domeneshop.no/admin?view=api\" target=\"_blank\">this page</a> after logging in to the control panel on our website: <a href=\"https://www.domeneshop.no/admin?view=api\" target=\"_blank\">https://www.domeneshop.no/admin?view=api</a> 
 *
 * API version: v0
 * Contact: kundeservice@domeneshop.no
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package domeneshop
// HttpForward struct for HttpForward
type HttpForward struct {
	// The subdomain this forward applies to, without the domain part.  For instance, `www` in the context of the `example.com` domain signifies a forward for `www.example.com`. 
	Host string `json:"host,omitempty"`
	// Whether to enable frame forwarding using an iframe embed. **NOT** recommended for a variety of reasons.
	Frame bool `json:"frame,omitempty"`
	// The URL to forward to. Must include scheme, e.g. `https://` or `ftp://`.
	Url string `json:"url,omitempty"`
}
