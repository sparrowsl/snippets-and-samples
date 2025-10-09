const std = @import("std");

var template =
    \\ <!DOCTYPE html>
    \\ <html lang="en">
    \\ <head>
    \\   <meta charset="UTF-8">
    \\   <meta name="viewport" content="width=device-width, initial-scale=1.0">
    \\   <meta name="author" content="{author}">
    \\   <title>{sitename}</title>
    \\ </head>
    \\ <body></body>
    \\ </html>
;

pub fn main() !void {
    var name_buffer: [100]u8 = undefined;
    var author_buffer: [100]u8 = undefined;
    var js_buffer: [2]u8 = undefined;
    var css_buffer: [2]u8 = undefined;
    var stdin = std.fs.File.stdin();

    std.debug.print("Site name: ", .{});
    var reader = stdin.reader(&name_buffer);
    const site_name = try reader.interface.takeDelimiterExclusive('\n');

    std.debug.print("Author: ", .{});
    reader = stdin.reader(&author_buffer);
    const author = try reader.interface.takeDelimiterExclusive('\n');

    std.debug.print("Do you want a folder for JavaScript? (y/n): ", .{});
    reader = stdin.reader(&js_buffer);
    const add_js = try reader.interface.takeByte();

    std.debug.print("Do you want a folder for CSS? (y/n): ", .{});
    reader = stdin.reader(&css_buffer);
    const add_css = try reader.interface.takeByte();

    std.fs.cwd().makeDir(site_name) catch |err| switch (err) {
        else => std.debug.print("Created ./{s}\n", .{site_name}),
    };

    const filename = try std.mem.concat(std.heap.page_allocator, u8, &[_][]const u8{ "./", site_name, "/index.html" });
    var buffer: [1024]u8 = undefined;

    var size = std.mem.replacementSize(u8, template, "{author}", author);
    _ = std.mem.replace(u8, template, "{author}", author, buffer[0..size]);
    const current = buffer[0..size];

    size = std.mem.replacementSize(u8, current, "{sitename}", site_name);
    _ = std.mem.replace(u8, current, "{sitename}", site_name, buffer[0..size]);

    _ = std.fs.cwd().writeFile(.{
        .sub_path = filename,
        .data = buffer[0..size],
        .flags = .{ .truncate = true },
    }) catch {};
    std.debug.print("Created ./{s}/index.html\n", .{site_name});

    if (add_js == 'y') {
        const parts = [_][]const u8{ "./", site_name, "/index.js" };
        _ = std.fs.cwd().writeFile(.{
            .sub_path = try std.mem.concat(std.heap.page_allocator, u8, &parts),
            .data = "",
            .flags = .{ .truncate = true },
        }) catch {};
        std.debug.print("Created ./{s}/index.js\n", .{site_name});
    }

    if (add_css == 'y') {
        const parts = [_][]const u8{ "./", site_name, "/index.css" };
        _ = std.fs.cwd().writeFile(.{
            .sub_path = try std.mem.concat(std.heap.page_allocator, u8, &parts),
            .data = "",
            .flags = .{ .truncate = true },
        }) catch {};
        std.debug.print("Created ./{s}/index.css\n", .{site_name});
    }
}
