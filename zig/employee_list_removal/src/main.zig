const std = @import("std");

pub fn main() !void {
    const stdin = std.io.getStdIn().reader();
    var gpa = std.heap.GeneralPurposeAllocator(.{}){};
    defer _ = gpa.deinit();
    const allocator = gpa.allocator();

    var file = std.fs.cwd().openFile("./employees.txt", .{ .mode = .read_only }) catch {
        std.debug.print("File not found or does not exists!!\n", .{});
        std.posix.exit(0);
    };
    defer file.close();

    var employees_list = std.ArrayList([]const u8).init(allocator);
    defer employees_list.deinit();

    var buffer: [100]u8 = undefined;
    const bytes_read = try file.read(&buffer);
    const buffer_slice = buffer[0..bytes_read];

    var tokens = std.mem.tokenizeSequence(u8, buffer_slice, "\n");
    while (tokens.next()) |token| {
        try employees_list.append(token);
    }

    display_employees(employees_list);

    std.debug.print("Enter an employee name to remove: ", .{});
    const name = try stdin.readUntilDelimiter(buffer[0..], '\n');

    remove_employee(&employees_list, name);

    display_employees(employees_list);
}

fn remove_employee(list: *std.ArrayList([]const u8), name: []const u8) void {
    for (list.items, 0..) |value, i| {
        if (std.mem.eql(u8, value, name)) {
            _ = list.orderedRemove(i);
            return;
        }
    }

    std.debug.print("The name \"{s}\" does not exists\n", .{name});
}

fn display_employees(employees: std.ArrayList([]const u8)) void {
    std.debug.print("There are {d} employees:\n", .{employees.items.len});
    for (employees.items) |employee| {
        std.debug.print("- {s}\n", .{employee});
    }
}
