import { useEffect } from "react";


export function useOutsideAlerter({ref, handleEvent}) {
  useEffect(() => {
    // Bind the event listener
    document.addEventListener("mousedown", handleEvent);
    return () => {
      // Unbind the event listener on clean up
      document.removeEventListener("mousedown", handleEvent);
    };
  }, [ref]);
}