# Scalable-E-Commerce-Platform
Trying to work on a project from roadmap sh https://roadmap.sh/projects/scalable-ecommerce-platform

## Microservices to implement

- User Service: Handles user registration, authentication, and profile management.
- Product Catalog Service: Manages product listings, categories, and inventory.
- Shopping Cart Service: Manages users’ shopping carts, including adding/removing items and updating quantities.
- Order Service: Processes orders, including placing orders, tracking order status, and managing order history.
- Payment Service: Handles payment processing, integrating with external payment gateways (e.g., Stripe, PayPal).
- Notification Service: Sends email and SMS notifications for various events (e.g., order confirmation, shipping updates). You can use third-party services like Twilio or SendGrid for this purpose.



-- Everything is implemented in gin

-- So, assuming the flow is as follows

User Service : User can create his profile and password
Cart Service : User can create cart and add items
Product Service: Something for admin, can create, delete, or modify products (such as inventory count, etc)
Order Service: When User clicks on Place Order on frontent, the request comes here, then it goes to product service to check if the products are available, and then allows to place order. After successful payment, stores in the database, and sends message to notification service
Notification Service: Receives message from Order Service and generates the notification appropriately