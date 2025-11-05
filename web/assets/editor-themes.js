(() => {
  function waitForMonaco() {
    return new Promise((resolve) => {
      const intervalId = setInterval(() => {
        if (typeof monaco !== "undefined") {
          clearInterval(intervalId);
          resolve();
        }
      }, 100);
    });
  }

  Promise.all([
    fetch("/assets/monaco-theme-light.json"),
    fetch("/assets/monaco-theme-dark.json")
  ])
    .then((themes) => Promise.all(
      themes.map((theme) => theme.json())
    ))
    .then(async ([lightTheme, darkTheme]) => {
      await waitForMonaco();
      monaco.editor.defineTheme("GithubLight", lightTheme);
      monaco.editor.defineTheme("GithubDark", darkTheme);
    })
})();
