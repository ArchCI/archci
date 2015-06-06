var gulp = require("gulp");
var buffer = require("vinyl-buffer");
var source = require("vinyl-source-stream");

var babelify = require("babelify");
var browserify = require("browserify");
var watchify = require("watchify");

var bundler = browserify({
    entries: ["./js/index.jsx"],
    transform: [babelify],
    debug: true,
});

function build() {
    var start = Date.now();
    console.log("[" + new Date().toLocaleTimeString() + "] Bundling...");

    bundler.bundle()
        .on("error", function(err) {
            console.error(err);
            this.emit("end");
        })
        .on("end", function() {
            console.log("[" + new Date().toLocaleTimeString() + "] Finished after " + (Date.now() - start) + " ms");
        })
        .pipe(source("bundle.js"))
        .pipe(buffer())
        .pipe(gulp.dest("./dist"));
}

gulp.task("watch", function() {
    bundler = watchify(bundler);
    bundler.on("update", function() { build() });
    build();
});

gulp.task("build", function() { build() });
gulp.task("default", ["watch"]);
