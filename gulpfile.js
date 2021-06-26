const gulp = require('gulp');
const uglify = require('gulp-uglify');
const terser = require('gulp-terser');
const sourcemaps = require('gulp-sourcemaps');
const proj = require("gulp-typescript").createProject("tsconfig.json");

// 全てのtsファイルから js と js.map を生成
gulp.task('default', build);
// srcフォルダの変更を監視してビルド
gulp.task('watch', function () {
  gulp.watch('./app/src/**/*.ts', build);
});

function build() {
  return proj.src()
    .pipe(sourcemaps.init())
    .pipe(proj()).js
    // ES2015以降の場合はuglifyをterserに変更
    .pipe(terser({
      // ステップ実行するためカンマ区切りで繋げない
      compress: {
        sequences: false
      }
    }))
    .pipe(sourcemaps.write('.', {
      sourceRoot: './',
      includeContent: false
    }))
    .pipe(gulp.dest("app/production/"));
}
