// Lint and build CSS
module.exports = function(grunt) {
  'use strict';

  // prettier-ignore
  grunt.registerTask('default', [
    'clean:build',
    'phantomjs',
  ]);

  // prettier-ignore
  grunt.registerTask('test', [
    'sasslint',
    'tslint',
    'typecheck',
    'exec:jest',
    'no-only-tests',
    'no-focus-convey-tests'
  ]);

  // prettier-ignore
  grunt.registerTask('tslint', [
    'newer:exec:tslintPackages',
    'newer:exec:tslintRoot'
  ]);

  // prettier-ignore
  grunt.registerTask('typecheck', [
    'newer:exec:typecheckPackages',
    'newer:exec:typecheckRoot'
  ]);

  grunt.registerTask('no-only-tests', function() {
    var files = grunt.file.expand(
      'public/**/*@(_specs|.test).@(ts|js|tsx|jsx)',
      'packages/smartems-data/**/*@(_specs|.test).@(ts|js|tsx|jsx)',
      'packages/**/*@(_specs|.test).@(ts|js|tsx|jsx)'
    );
    grepFiles(files, '.only(', 'found only statement in test: ');
  });

  grunt.registerTask('no-focus-convey-tests', function() {
    var files = grunt.file.expand('pkg/**/*_test.go');
    grepFiles(files, 'FocusConvey(', 'found FocusConvey statement in test: ');
  });

  function grepFiles(files, pattern, errorMessage) {
    files.forEach(function(spec) {
      var rows = grunt.file.read(spec).split('\n');
      rows.forEach(function(row) {
        if (row.indexOf(pattern) > 0) {
          grunt.log.errorlns(row);
          grunt.fail.warn(errorMessage + spec);
        }
      });
    });
  }
};
