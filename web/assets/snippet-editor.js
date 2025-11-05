(() => {
  const editorSymbol = Symbol.for("ActiveEditor");
  const nameRegex = /^(#|\/\/|\/\*|\/\*\*)\s?filename:(.*)$/i;
  const languages = {
    js: "javascript",
    ts: "typescript",
    sh: "shell",
    json: "json",
  };

  function getLanguage(filename, fallback = "shell") {
    const ext = filename.split(".").pop();
    return languages[ext] || ext || fallback;
  }

  function createFilenameLine(name) {
    const language = getLanguage(name);
    const comment = language === "shell" ? "#" : "//";
    return `${comment} filename: ${name}`;
  }

  function setEditorTheme(theme) {
    const editorTheme = {
      dark: "GithubDark",
      light: "GithubLight",
      emerald: "GithubLight",
    }[theme];

    monaco.editor.setTheme(editorTheme);
  }

  function getFormSubmit() {
    return document
      .getElementById("snippet-form")
      ?.querySelector("button[type=submit]");
  }

  function handleSave() {
    getFormSubmit()?.click();
  }

  const container = document.getElementById("container");
  const contents = document.getElementById("contents");
  const filename = document.getElementById("filename");
  const filenameDisplay = document.getElementById("filename-display");
  container.innerHTML = "";

  if (!filename.value) {
    filename.value = "snippet.sh";
    filenameDisplay.textContent = filename.value;
  }

  document.title = `${filename.value} - Snippet Editor`;

  window[editorSymbol]?.dispose();
  const editor = monaco.editor.create(container, {
    automaticLayout: true,
    bracketPairColorization: {
      enabled: true,
    },
    cursorBlinking: "expand",
    detectIndentation: false,
    fontFamily: "JetBrains Mono",
    fontLigatures: true,
    fontSize: 18,
    insertSpaces: true,
    language: getLanguage(filename.value),
    lineHeight: 1.4,
    padding: {
      top: 16,
      bottom: 16,
    },
    scrollBeyondLastLine: false,
    suggest: {
      showFields: false,
      showFunctions: false,
    },
    tabSize: 2,
    theme: "GithubDark",
    value: `${createFilenameLine(filename.value)}\n${contents.value}`,
    wordWrap: "on",
  });
  window[editorSymbol] = editor;

  window
    .matchMedia("(prefers-color-scheme: dark)")
    .addEventListener("change", (event) => {
      const documentTheme = document.documentElement.getAttribute("data-theme");
      const theme = documentTheme
        ? documentTheme
        : event.matches
        ? "dark"
        : "light";

      setEditorTheme(theme);
    });

  window.addEventListener("theme-change", (event) => {
    const { theme } = event.detail;
    setEditorTheme(theme);
  });

  setEditorTheme(document.documentElement.getAttribute("data-theme"));

  editor.addCommand(monaco.KeyMod.CtrlCmd | monaco.KeyCode.KeyS, handleSave);

  editor.onDidChangeModelContent(() => {
    const value = editor.getValue();
    const lines = value.split("\n");
    const firstLine = lines[0] || "";
    const nameMatches = firstLine.match(nameRegex);
    const name = nameMatches?.[2]?.trim() || "";
    const language = getLanguage(name);
    const contentsValue = lines.slice(1).join("\n");

    contents.value = contentsValue.trim();
    filename.value = name;
    filenameDisplay.textContent = name;

    monaco.editor.setModelLanguage(editor.getModel(), language);

    const filenameLine = createFilenameLine(name);
    if (lines[0] !== filenameLine) {
      document.title = `${filename.value} - Snippet Editor`;
      editor.executeEdits("snippet-editor", [
        {
          range: editor.getModel().getFullModelRange(),
          text: [filenameLine, ...lines.slice(1)].join("\n"),
        },
      ]);
    }

    if (contents.value) {
      getFormSubmit()?.removeAttribute("disabled");
    } else {
      getFormSubmit()?.setAttribute("disabled", "");
    }
  });

  editor.focus();
})();
