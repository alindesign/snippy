document.addEventListener("DOMContentLoaded", () => {
  function getTheme() {
    const isDark = document.querySelector(".theme-controller").checked;
    const theme = isDark ? "dark" : "emerald";

    return { theme, isDark, isLight: !isDark };
  }

  // Active navigation links
  const navPages = document.querySelectorAll(".btn-nav");
  navPages.forEach((navPage) => {
    const url = new URL(navPage.href ?? "/");
    const active = url.pathname && url.pathname === window.location.pathname;

    navPage.classList.toggle("btn-neutral", active);
    navPage.classList.toggle("btn-ghost", !active);
  });

  // initialize theme
  const theme =
    localStorage.getItem("theme") ||
    window.matchMedia("(prefers-color-scheme: dark)").matches
      ? "dark"
      : "emerald";

  document.documentElement.setAttribute("data-theme", theme);
  document.querySelector(".theme-controller").checked = theme === "dark";
  window.dispatchEvent(new CustomEvent("theme-change", { detail: getTheme() }));

  document
    .querySelector(".theme-controller")
    ?.addEventListener("change", () => {
      const theme = getTheme();
      document.documentElement.setAttribute("data-theme", theme.theme);
      localStorage.setItem("theme", theme.theme);
      window.dispatchEvent(new CustomEvent("theme-change", { detail: theme }));
    });
});
