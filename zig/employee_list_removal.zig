const std = @import("std");

pub fn main() !void {
    const allocator = std.heap.page_allocator;

    const stdin = std.io.getStdIn().reader();
    var buffer: [1024]u8 = undefined;

    var employees_list = std.ArrayList([]const u8).init(allocator);
    defer employees_list.deinit();
    try employees_list.appendSlice(&[_][]const u8{
        "John Smith",
        "Jackie Jackson",
        "Chris Jones",
        "Amanda Cullen",
        "Jeremy Goodwin",
    });

    display_employees(&employees_list);

    std.debug.print("Enter an employee name to remove: ", .{});
    const name = try stdin.readUntilDelimiter(buffer[0..], '\n');

    remove_employee(&employees_list, name);

    display_employees(&employees_list);
}

fn remove_employee(list: *std.ArrayList([]const u8), name: []const u8) void {
    for (list.items, 0..) |value, i| {
        if (std.mem.eql(u8, value, name)) {
            _ = list.orderedRemove(i);
            break;
        }
    }

    std.debug.print("The name \"{s}\" does not exists", .{name});
}

fn display_employees(list: *std.ArrayList([]const u8)) void {
    std.debug.print("There are {d} employees:\n", .{list.items.len});
    for (list.items) |employee| {
        std.debug.print("- {s}\n", .{employee});
    }
}
