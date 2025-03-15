const std = @import("std");

pub fn main() !void {
    const stdin = std.io.getStdIn().reader();
    var buffer: [20]u8 = undefined;

    const Item = struct { price: f64, quantity: i32 };

    var items: [3]Item = undefined;

    for (0..items.len) |idx| {
        std.debug.print("Enter the price of item {d}: ", .{idx + 1});
        var input = try stdin.readUntilDelimiter(buffer[0..], '\n');
        const price = try std.fmt.parseFloat(f64, input);

        std.debug.print("Enter the quantity of item {d}: ", .{idx + 1});
        input = try stdin.readUntilDelimiter(buffer[0..], '\n');
        const quantity = try std.fmt.parseInt(i32, input, 10);

        items[idx] = Item{ .price = price, .quantity = quantity };
    }

    var subtotal: f64 = 0;
    for (0..items.len) |idx| {
        const tmp: f64 = @floatFromInt(items[idx].quantity);
        subtotal = subtotal + (items[idx].price * tmp);
    }

    const tax_rate: f64 = 5.5;
    const tax: f64 = (tax_rate / 100) * subtotal;

    std.debug.print("Subtotal: ${d:.2}\n", .{subtotal});
    std.debug.print("Tax: ${d:.2}\n", .{tax});
    std.debug.print("Total: ${d:.2}\n", .{subtotal + tax});
}
