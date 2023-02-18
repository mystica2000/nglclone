import { createSignal } from "solid-js";
import CopyToClipboardIcon from "../assets/icons/copy";
import { Show } from "solid-js";
import MarkIcon from "../assets/icons/mark";

export default function CopyToClipBoard(props: {link:string}) {

  const [copy,setCopy] = createSignal(false);

  const handleCopy = async () => {
    try {
      await navigator.clipboard.writeText(props.link);
      setCopy(true);
      setTimeout(()=> {
        setCopy(false);
      },3000);
    } catch(e) {
      console.error(e);
    }
  }

  return (
    <button onClick={handleCopy}>
      <Show when={copy()}
      fallback={<><CopyToClipboardIcon /> Copy Link</>}>
      <><MarkIcon /> Copied! </>
      </Show>
    </button>

  )
}