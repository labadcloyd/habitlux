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

export function calcBgColor(ratio, color) {
  if(isNaN(ratio)) { return `rgb(${color} / 0%)` }
  if (ratio >= 1) { return `rgb(${color} / 100%)` }
  if (ratio >= 0.5) { return `rgb(${color} / 60%)` }
  if (ratio >= 0.1) { return `rgb(${color} / 30%)` }
  if (ratio < 0.1) { return `rgb(${color} / 0%)` }
}