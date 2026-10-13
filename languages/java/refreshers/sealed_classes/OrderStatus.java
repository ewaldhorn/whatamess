class Instant {}

public sealed interface OrderStatus permits OrderStatus.Pending, OrderStatus.Shipped, OrderStatus.Delivered, OrderStatus.Cancelled {
    record Pending(Instant createdTime) implements OrderStatus {}
    record Shipped(String trackingNumber, Instant shippedTime) implements OrderStatus {}
    record Delivered(Instant deliveryTime, String recipientSignature) implements OrderStatus {}
    record Cancelled(String reason, Instant cancelTime) implements OrderStatus {}
}
