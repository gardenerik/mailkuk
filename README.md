<div align="center">
  <h3>Mailkuk</h3>
  <p>Simple SMTP server that sends incoming mail to an HTTP endpoint.</p>
</div>

## About

`mailkuk` is a basic SMTP submission server. Its sole purpose is to receive email over SMTP and send it to an HTTP endpoint.

Useful if you want to programatically handle incoming email (notifications from external services, servers, printers...).

## Usage

We only have a single configuration file `config.yml`. An example configuration with additional comments is provided below.

```yaml
server:
  # Which address should the SMTP server listen on.
  listen_addr: "0.0.0.0:25"

  # Domain is displayed in the SMTP greeting. 
  domain: "mailkuk.example.com"
  
  # Print debug data?
  debug: false
  
# Email routing configuration
routing:
    # All email sent to this adress...
  - mail: incoming@example.com
    # ...will be sent to this HTTP endpoint...
    url: https://example.com/mail_handler/
    # ...with these (optional) headers.
    headers:
      Authorization: Bearer secret_token
```

### HTTP request

**Warning:** `mailkuk` sends the whole email as received without doing any parsing / sanitizing – make sure to handle it accordingly.

The whole email data (a `message/rfc822` encapsulated message) gets sent as the request body using a POST request, with the configured headers.

Email body is limited to 5 MB.
