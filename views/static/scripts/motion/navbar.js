import { animate, stagger } from "/static/scripts/motion.min.js";

const navMenu = document.getElementById("navMenu");
const expandedNavMenu = document.getElementById("expandedNavMenu");
const iconP1 = document.getElementById("icon-p1");
const iconP2 = document.getElementById("icon-p2");
const iconP3 = document.getElementById("icon-p3");

const iconPaths = {
  avatar: {
    p1: "M 9,8 C 9,0, 23,0, 23,8",
    p2: "M 9,8 C 9,16, 23,16, 23,8",
    p3: "M 6,26 C 6,18, 26,18, 26,26",
  },
  hamburger: {
    p1: "M 4,10 C 12,10 20,10 28,10",
    p2: "M 4,16 C 12,16 20,16 28,16",
    p3: "M 4,22 C 12,22 20,22 28,22",
  },
  close: {
    p1: "M 8,8 C 12,12 20,20 24,24",
    p2: "M 16,16 C 16,16 16,16 16,16",
    p3: "M 8,24 C 12,20 20,12 24,8",
  },
};

function setIcon(paths, options = {}) {
  const t = options.transition ?? { duration: 0.3, ease: "easeOut" };
  animate(iconP1, { d: paths.p1 }, t);
  animate(iconP2, { d: paths.p2 }, t);
  animate(iconP3, { d: paths.p3 }, t);
}

function navbarScrolling(scrolling, quick = false) {
  const bookmarks = document.querySelectorAll(".bookmark_item");
  const telescopeTrigger = document.querySelector("telescopeTrigger");

  if (scrolling) {
    animate(
      bookmarks,
      { width: 0, opacity: 0, marginLeft: 0 },
      { duration: quick ? 0.25 : 0.75, delay: stagger(quick ? 0.055 : 0.087, { from: "first" }), ease: "anticipate" }
    );
    animate(telescopeTrigger, { x: "6.75rem" }, { duration: quick ? 0.37 : 1, ease: "anticipate" });
    if (!navMenu.classList.contains("menu-open")) {
      setIcon(iconPaths.hamburger);
    }
  }

  if (!scrolling) {
    animate(
      bookmarks,
      { width: "auto", opacity: 1, marginLeft: "0.75rem" },
      { duration: 0.75, delay: stagger(0.037, { from: "last" }), ease: "backInOut" }
    );
    if (bookmarks.length > 0) {
      animate(bookmarks[0], { marginLeft: 0 });
    }
    animate(telescopeTrigger, { x: "0" }, { duration: 1, ease: "backInOut" });
    if (!navMenu.classList.contains("menu-open")) {
      setIcon(iconPaths.avatar);
    }
  }
}

function closeMenu() {
  if (navMenu.classList.contains("menu-open")) {
    navMenu.classList.remove("menu-open");
    toggleMenu();
  }
}

function toggleMenu() {
  const isOpen = navMenu.classList.contains("menu-open");
  const navbar = document.querySelector("navbar");

  const menuBackdrop = document.getElementById("menu-backdrop");
  const menuBg = document.getElementById("menu-bg");
  const navLinks = menuBg?.querySelectorAll(".page-link");
  const accountLinks = menuBg?.querySelectorAll(".account-link");

  if (isOpen) {
    navbar.classList.remove("navbar_scrolling");
    if (window.scrollY == 0) {
      navbarScrolling(true, true);
    }

    setIcon(iconPaths.close);
    expandedNavMenu.style.pointerEvents = "auto";

    animate(menuBackdrop, { opacity: 1 }, { duration: 0.3 });
    animate(menuBg, { transform: "translateX(0%)" }, { type: "spring", bounce: 0.4, duration: 0.7 });
    if (navLinks) {
      animate(
        navLinks,
        { opacity: [0, 1], x: [20, 0] },
        {
          ease: "backInOut",
          duration: 0.3,
          delay: stagger(0.1),
        }
      );
    }
    if (accountLinks) {
      animate(
        accountLinks,
        { opacity: [0, 1], x: [20, 0] },
        {
          ease: "backInOut",
          duration: 0.3,
          delay: stagger(0.1, { from: "last" }),
        }
      );
    }
  }

  if (!isOpen) {
    if (window.scrollY == 0) {
      navbarScrolling(false);
    } else {
      navbar.classList.add("navbar_scrolling");
    }
    setIcon(window.scrollY > 0 ? iconPaths.hamburger : iconPaths.avatar);
    expandedNavMenu.style.pointerEvents = "none";

    animate(navMenu, { rotate: 0 }, { ease: "easeOut", duration: 0.3 });
    animate(menuBackdrop, { opacity: 0 }, { duration: 0.3, delay: 0.1 });
    animate(menuBg, { transform: "translateX(100%)" }, { type: "spring" });
    if (navLinks) {
      animate(navLinks, { opacity: 0, x: 20 }, { ease: "easeOut", duration: 0.2, delay: stagger(0.08) });
    }
  }
}

window.navbarScrolling = navbarScrolling;
window.toggleMenu = toggleMenu;
window.closeMenu = closeMenu;
