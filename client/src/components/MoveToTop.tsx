import { createEffect, createSignal, Show } from "solid-js";
import TopArrowIcon from "../assets/icons/top"

export default function MoveToTop() {

  const [scrollToTop, setScrollToTop] = createSignal(false);

  createEffect(() => {
    window.addEventListener("scroll", () => {
      if (window.scrollY > 100) {
        setScrollToTop(true);
      } else {
        setScrollToTop(false);
      }
    })
  });

  const scrollUp = () => {
    window.scrollTo({
      top: 0,
      behavior: "smooth"
    })
  }

  return (
    <Show
      when={scrollToTop()}
      fallback={""}
    >
      <button onClick={scrollUp} class="scrollToTopBtn">
        <TopArrowIcon />
      </button>
    </Show>
  )
}