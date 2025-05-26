# Orders

Params Types:

- <a href="https://pkg.go.dev/github.com/ankitdas13/gotue">gotue</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/gotue#NotesUnionParam">NotesUnionParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/ankitdas13/gotue">gotue</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/gotue#NotesUnion">NotesUnion</a>
- <a href="https://pkg.go.dev/github.com/ankitdas13/gotue">gotue</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/gotue#Order">Order</a>
- <a href="https://pkg.go.dev/github.com/ankitdas13/gotue">gotue</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/gotue#OrderListResponse">OrderListResponse</a>
- <a href="https://pkg.go.dev/github.com/ankitdas13/gotue">gotue</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/gotue#OrderListPaymentsResponse">OrderListPaymentsResponse</a>

Methods:

- <code title="post /orders">client.Orders.<a href="https://pkg.go.dev/github.com/ankitdas13/gotue#OrderService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/ankitdas13/gotue">gotue</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/gotue#OrderNewParams">OrderNewParams</a>) (<a href="https://pkg.go.dev/github.com/ankitdas13/gotue">gotue</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/gotue#Order">Order</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /orders/{id}">client.Orders.<a href="https://pkg.go.dev/github.com/ankitdas13/gotue#OrderService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/ankitdas13/gotue">gotue</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/gotue#Order">Order</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /orders/{id}">client.Orders.<a href="https://pkg.go.dev/github.com/ankitdas13/gotue#OrderService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/ankitdas13/gotue">gotue</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/gotue#OrderUpdateParams">OrderUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/ankitdas13/gotue">gotue</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/gotue#Order">Order</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /orders">client.Orders.<a href="https://pkg.go.dev/github.com/ankitdas13/gotue#OrderService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/ankitdas13/gotue">gotue</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/gotue#OrderListParams">OrderListParams</a>) (<a href="https://pkg.go.dev/github.com/ankitdas13/gotue">gotue</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/gotue#OrderListResponse">OrderListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /orders/{id}/payments">client.Orders.<a href="https://pkg.go.dev/github.com/ankitdas13/gotue#OrderService.ListPayments">ListPayments</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/ankitdas13/gotue">gotue</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/gotue#OrderListPaymentsParams">OrderListPaymentsParams</a>) (<a href="https://pkg.go.dev/github.com/ankitdas13/gotue">gotue</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/gotue#OrderListPaymentsResponse">OrderListPaymentsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Customers

Response Types:

- <a href="https://pkg.go.dev/github.com/ankitdas13/gotue">gotue</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/gotue#Customer">Customer</a>
- <a href="https://pkg.go.dev/github.com/ankitdas13/gotue">gotue</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/gotue#CustomerListResponse">CustomerListResponse</a>

Methods:

- <code title="post /customers">client.Customers.<a href="https://pkg.go.dev/github.com/ankitdas13/gotue#CustomerService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/ankitdas13/gotue">gotue</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/gotue#CustomerNewParams">CustomerNewParams</a>) (<a href="https://pkg.go.dev/github.com/ankitdas13/gotue">gotue</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/gotue#Customer">Customer</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /customers/{id}">client.Customers.<a href="https://pkg.go.dev/github.com/ankitdas13/gotue#CustomerService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/ankitdas13/gotue">gotue</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/gotue#Customer">Customer</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /customers/{id}">client.Customers.<a href="https://pkg.go.dev/github.com/ankitdas13/gotue#CustomerService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/ankitdas13/gotue">gotue</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/gotue#CustomerUpdateParams">CustomerUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/ankitdas13/gotue">gotue</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/gotue#Customer">Customer</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /customers">client.Customers.<a href="https://pkg.go.dev/github.com/ankitdas13/gotue#CustomerService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/ankitdas13/gotue">gotue</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/gotue#CustomerListParams">CustomerListParams</a>) (<a href="https://pkg.go.dev/github.com/ankitdas13/gotue">gotue</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/gotue#CustomerListResponse">CustomerListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Payments

Response Types:

- <a href="https://pkg.go.dev/github.com/ankitdas13/gotue">gotue</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/gotue#AcquirerData">AcquirerData</a>
- <a href="https://pkg.go.dev/github.com/ankitdas13/gotue">gotue</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/gotue#Card">Card</a>
- <a href="https://pkg.go.dev/github.com/ankitdas13/gotue">gotue</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/gotue#Payment">Payment</a>
- <a href="https://pkg.go.dev/github.com/ankitdas13/gotue">gotue</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/gotue#Refund">Refund</a>
- <a href="https://pkg.go.dev/github.com/ankitdas13/gotue">gotue</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/gotue#PaymentListResponse">PaymentListResponse</a>

Methods:

- <code title="get /payments/{id}">client.Payments.<a href="https://pkg.go.dev/github.com/ankitdas13/gotue#PaymentService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/ankitdas13/gotue">gotue</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/gotue#PaymentGetParams">PaymentGetParams</a>) (<a href="https://pkg.go.dev/github.com/ankitdas13/gotue">gotue</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/gotue#Payment">Payment</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /payments/{id}">client.Payments.<a href="https://pkg.go.dev/github.com/ankitdas13/gotue#PaymentService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/ankitdas13/gotue">gotue</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/gotue#PaymentUpdateParams">PaymentUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/ankitdas13/gotue">gotue</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/gotue#Payment">Payment</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /payments">client.Payments.<a href="https://pkg.go.dev/github.com/ankitdas13/gotue#PaymentService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/ankitdas13/gotue">gotue</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/gotue#PaymentListParams">PaymentListParams</a>) (<a href="https://pkg.go.dev/github.com/ankitdas13/gotue">gotue</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/gotue#PaymentListResponse">PaymentListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /payments/{id}/capture">client.Payments.<a href="https://pkg.go.dev/github.com/ankitdas13/gotue#PaymentService.Capture">Capture</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/ankitdas13/gotue">gotue</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/gotue#PaymentCaptureParams">PaymentCaptureParams</a>) (<a href="https://pkg.go.dev/github.com/ankitdas13/gotue">gotue</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/gotue#Payment">Payment</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## CardDetails

Methods:

- <code title="get /payments/{id}/card">client.Payments.CardDetails.<a href="https://pkg.go.dev/github.com/ankitdas13/gotue#PaymentCardDetailService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/ankitdas13/gotue">gotue</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/gotue#Card">Card</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## QrCodes

Response Types:

- <a href="https://pkg.go.dev/github.com/ankitdas13/gotue">gotue</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/gotue#QrCode">QrCode</a>
- <a href="https://pkg.go.dev/github.com/ankitdas13/gotue">gotue</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/gotue#PaymentQrCodeListResponse">PaymentQrCodeListResponse</a>

Methods:

- <code title="post /payments/qr_codes">client.Payments.QrCodes.<a href="https://pkg.go.dev/github.com/ankitdas13/gotue#PaymentQrCodeService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/ankitdas13/gotue">gotue</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/gotue#PaymentQrCodeNewParams">PaymentQrCodeNewParams</a>) (<a href="https://pkg.go.dev/github.com/ankitdas13/gotue">gotue</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/gotue#QrCode">QrCode</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /payments/qr_codes/{id}">client.Payments.QrCodes.<a href="https://pkg.go.dev/github.com/ankitdas13/gotue#PaymentQrCodeService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/ankitdas13/gotue">gotue</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/gotue#PaymentQrCodeGetParams">PaymentQrCodeGetParams</a>) (<a href="https://pkg.go.dev/github.com/ankitdas13/gotue">gotue</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/gotue#QrCode">QrCode</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /payments/qr_codes">client.Payments.QrCodes.<a href="https://pkg.go.dev/github.com/ankitdas13/gotue#PaymentQrCodeService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/ankitdas13/gotue">gotue</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/gotue#PaymentQrCodeListParams">PaymentQrCodeListParams</a>) (<a href="https://pkg.go.dev/github.com/ankitdas13/gotue">gotue</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/gotue#PaymentQrCodeListResponse">PaymentQrCodeListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /payments/qr_codes/{id}/close">client.Payments.QrCodes.<a href="https://pkg.go.dev/github.com/ankitdas13/gotue#PaymentQrCodeService.Close">Close</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/ankitdas13/gotue">gotue</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/gotue#QrCode">QrCode</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Refunds

Response Types:

- <a href="https://pkg.go.dev/github.com/ankitdas13/gotue">gotue</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/gotue#RefundList">RefundList</a>

Methods:

- <code title="post /payments/{id}/refund">client.Payments.Refunds.<a href="https://pkg.go.dev/github.com/ankitdas13/gotue#PaymentRefundService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/ankitdas13/gotue">gotue</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/gotue#PaymentRefundNewParams">PaymentRefundNewParams</a>) (<a href="https://pkg.go.dev/github.com/ankitdas13/gotue">gotue</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/gotue#Refund">Refund</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /payments/{id}/refunds/{refund_id}">client.Payments.Refunds.<a href="https://pkg.go.dev/github.com/ankitdas13/gotue#PaymentRefundService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, refundID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/ankitdas13/gotue">gotue</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/gotue#PaymentRefundGetParams">PaymentRefundGetParams</a>) (<a href="https://pkg.go.dev/github.com/ankitdas13/gotue">gotue</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/gotue#Refund">Refund</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /payments/{id}/refunds">client.Payments.Refunds.<a href="https://pkg.go.dev/github.com/ankitdas13/gotue#PaymentRefundService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/ankitdas13/gotue">gotue</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/gotue#PaymentRefundListParams">PaymentRefundListParams</a>) (<a href="https://pkg.go.dev/github.com/ankitdas13/gotue">gotue</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/gotue#RefundList">RefundList</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# PaymentLinks

Response Types:

- <a href="https://pkg.go.dev/github.com/ankitdas13/gotue">gotue</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/gotue#PaymentLink">PaymentLink</a>
- <a href="https://pkg.go.dev/github.com/ankitdas13/gotue">gotue</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/gotue#PaymentLinkListResponse">PaymentLinkListResponse</a>
- <a href="https://pkg.go.dev/github.com/ankitdas13/gotue">gotue</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/gotue#PaymentLinkNotifyResponse">PaymentLinkNotifyResponse</a>

Methods:

- <code title="post /payment_links">client.PaymentLinks.<a href="https://pkg.go.dev/github.com/ankitdas13/gotue#PaymentLinkService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/ankitdas13/gotue">gotue</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/gotue#PaymentLinkNewParams">PaymentLinkNewParams</a>) (<a href="https://pkg.go.dev/github.com/ankitdas13/gotue">gotue</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/gotue#PaymentLink">PaymentLink</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /payment_links/{id}">client.PaymentLinks.<a href="https://pkg.go.dev/github.com/ankitdas13/gotue#PaymentLinkService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/ankitdas13/gotue">gotue</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/gotue#PaymentLink">PaymentLink</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /payment_links/{id}">client.PaymentLinks.<a href="https://pkg.go.dev/github.com/ankitdas13/gotue#PaymentLinkService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/ankitdas13/gotue">gotue</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/gotue#PaymentLinkUpdateParams">PaymentLinkUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/ankitdas13/gotue">gotue</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/gotue#PaymentLink">PaymentLink</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /payment_links">client.PaymentLinks.<a href="https://pkg.go.dev/github.com/ankitdas13/gotue#PaymentLinkService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/ankitdas13/gotue">gotue</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/gotue#PaymentLinkListParams">PaymentLinkListParams</a>) (<a href="https://pkg.go.dev/github.com/ankitdas13/gotue">gotue</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/gotue#PaymentLinkListResponse">PaymentLinkListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /payment_links/{id}/notify_by/{medium}">client.PaymentLinks.<a href="https://pkg.go.dev/github.com/ankitdas13/gotue#PaymentLinkService.Notify">Notify</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, medium <a href="https://pkg.go.dev/github.com/ankitdas13/gotue">gotue</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/gotue#PaymentLinkNotifyParamsMedium">PaymentLinkNotifyParamsMedium</a>, body <a href="https://pkg.go.dev/github.com/ankitdas13/gotue">gotue</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/gotue#PaymentLinkNotifyParams">PaymentLinkNotifyParams</a>) (<a href="https://pkg.go.dev/github.com/ankitdas13/gotue">gotue</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/gotue#PaymentLinkNotifyResponse">PaymentLinkNotifyResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Refunds

Methods:

- <code title="get /refunds/{id}">client.Refunds.<a href="https://pkg.go.dev/github.com/ankitdas13/gotue#RefundService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/ankitdas13/gotue">gotue</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/gotue#Refund">Refund</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /refunds">client.Refunds.<a href="https://pkg.go.dev/github.com/ankitdas13/gotue#RefundService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/ankitdas13/gotue">gotue</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/gotue#RefundListParams">RefundListParams</a>) (<a href="https://pkg.go.dev/github.com/ankitdas13/gotue">gotue</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/gotue#RefundList">RefundList</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /refunds/{id}">client.Refunds.<a href="https://pkg.go.dev/github.com/ankitdas13/gotue#RefundService.UpdateNotes">UpdateNotes</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/ankitdas13/gotue">gotue</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/gotue#RefundUpdateNotesParams">RefundUpdateNotesParams</a>) (<a href="https://pkg.go.dev/github.com/ankitdas13/gotue">gotue</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/gotue#Refund">Refund</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Settlements

Response Types:

- <a href="https://pkg.go.dev/github.com/ankitdas13/gotue">gotue</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/gotue#Settlement">Settlement</a>
- <a href="https://pkg.go.dev/github.com/ankitdas13/gotue">gotue</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/gotue#SettlementListResponse">SettlementListResponse</a>

Methods:

- <code title="get /settlements/{id}">client.Settlements.<a href="https://pkg.go.dev/github.com/ankitdas13/gotue#SettlementService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/ankitdas13/gotue">gotue</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/gotue#Settlement">Settlement</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /settlements">client.Settlements.<a href="https://pkg.go.dev/github.com/ankitdas13/gotue#SettlementService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/ankitdas13/gotue">gotue</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/gotue#SettlementListParams">SettlementListParams</a>) (<a href="https://pkg.go.dev/github.com/ankitdas13/gotue">gotue</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/gotue#SettlementListResponse">SettlementListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Recon

Response Types:

- <a href="https://pkg.go.dev/github.com/ankitdas13/gotue">gotue</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/gotue#SettlementReconGetResponse">SettlementReconGetResponse</a>

Methods:

- <code title="get /settlements/recon/combined">client.Settlements.Recon.<a href="https://pkg.go.dev/github.com/ankitdas13/gotue#SettlementReconService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/ankitdas13/gotue">gotue</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/gotue#SettlementReconGetParams">SettlementReconGetParams</a>) (<a href="https://pkg.go.dev/github.com/ankitdas13/gotue">gotue</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/gotue#SettlementReconGetResponse">SettlementReconGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Ondemand

Response Types:

- <a href="https://pkg.go.dev/github.com/ankitdas13/gotue">gotue</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/gotue#SettlementOndemand">SettlementOndemand</a>

Methods:

- <code title="post /settlements/ondemand">client.Settlements.Ondemand.<a href="https://pkg.go.dev/github.com/ankitdas13/gotue#SettlementOndemandService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/ankitdas13/gotue">gotue</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/gotue#SettlementOndemandNewParams">SettlementOndemandNewParams</a>) (<a href="https://pkg.go.dev/github.com/ankitdas13/gotue">gotue</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/gotue#SettlementOndemand">SettlementOndemand</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /settlements/ondemand/{id}">client.Settlements.Ondemand.<a href="https://pkg.go.dev/github.com/ankitdas13/gotue#SettlementOndemandService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/ankitdas13/gotue">gotue</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/gotue#SettlementOndemand">SettlementOndemand</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Payouts

Response Types:

- <a href="https://pkg.go.dev/github.com/ankitdas13/gotue">gotue</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/gotue#Payout">Payout</a>
- <a href="https://pkg.go.dev/github.com/ankitdas13/gotue">gotue</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/gotue#PayoutListResponse">PayoutListResponse</a>

Methods:

- <code title="get /payouts/{id}">client.Payouts.<a href="https://pkg.go.dev/github.com/ankitdas13/gotue#PayoutService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/ankitdas13/gotue">gotue</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/gotue#Payout">Payout</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /payouts">client.Payouts.<a href="https://pkg.go.dev/github.com/ankitdas13/gotue#PayoutService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/ankitdas13/gotue">gotue</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/gotue#PayoutListParams">PayoutListParams</a>) (<a href="https://pkg.go.dev/github.com/ankitdas13/gotue">gotue</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/gotue#PayoutListResponse">PayoutListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Plans

Response Types:

- <a href="https://pkg.go.dev/github.com/ankitdas13/gotue">gotue</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/gotue#Plan">Plan</a>

Methods:

- <code title="post /plans">client.Plans.<a href="https://pkg.go.dev/github.com/ankitdas13/gotue#PlanService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/ankitdas13/gotue">gotue</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/gotue#PlanNewParams">PlanNewParams</a>) (<a href="https://pkg.go.dev/github.com/ankitdas13/gotue">gotue</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/gotue#Plan">Plan</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
