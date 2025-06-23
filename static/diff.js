require.config({ paths: { vs: 'https://cdn.jsdelivr.net/npm/monaco-editor@0.41.0/min/vs' }});

require(['vs/editor/editor.main'], function () {
  const originalText = `Hello world
This is a test
Line three
Line four`;

  const modifiedText = `Hello world
This is a fest
Line 3
Line four`;

  const originalModel = monaco.editor.createModel(originalText, 'plaintext');
  const modifiedModel = monaco.editor.createModel(modifiedText, 'plaintext');

  const diffEditor = monaco.editor.createDiffEditor(document.getElementById('container'), {
    theme: 'vs-dark',
    enableSplitViewResizing: true,
    renderSideBySide: true,
    automaticLayout: true,
  });

  diffEditor.setModel({
    original: originalModel,
    modified: modifiedModel
  });
});
