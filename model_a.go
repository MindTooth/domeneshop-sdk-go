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
// A struct for A
type A struct {
	// ID of DNS record
	Id int32 `json:"id"`
	// The host/subdomain the DNS record applies to
	Host string `json:"host"`
	// TTL of DNS record in seconds. Must be a multiple of 60.
	Ttl int32 `json:"ttl,omitempty"`
	Type string `json:"type"`
	// IPv4 address
	Data string `json:"data"`
}
