(() => {
  fetch("/assets/monaco-theme-light.json")
    .then((response) => response.json())
    .then((theme) => {
      monaco.editor.defineTheme("GithubLight", theme);
    });

  fetch("/assets/monaco-theme-dark.json")
    .then((response) => response.json())
    .then((theme) => {
      monaco.editor.defineTheme("GithubDark", theme);
    });
})();
