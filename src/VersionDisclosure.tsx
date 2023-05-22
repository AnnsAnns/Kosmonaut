import { Disclosure, Transition } from "@headlessui/react";
import { useState, useEffect } from "react";
import axios from "axios";

export default function MyDisclosure() {
  const [post, setPost] = useState("");
  useEffect(() => {
    axios
      .get(
        "https://raw.githubusercontent.com/tumGER/Kosmonaut/gh-pages/assets/version.txt"
      )
      .catch(() => {
        setPost("Error\nServer returned 404!");
      })
      .then((response) => {
        if (response && response.data) {
            setPost(response.data);
        } else {
            setPost("Error\nReceived incorrect response!")
        }
      });
  }, []);
  return (
    <Disclosure>
      <Disclosure.Button className="text-lg font-light">
        Click to see version
      </Disclosure.Button>
      <Transition
        enter="transition duration-100 ease-out"
        enterFrom="transform scale-95 opacity-0"
        enterTo="transform scale-100 opacity-100"
        leave="transition duration-75 ease-out"
        leaveFrom="transform scale-100 opacity-100"
        leaveTo="transform scale-95 opacity-0"
      >
        <Disclosure.Panel className="text-gray-500">
          {post != null ? (
            <div className="whitespace-pre-line">{post}</div>
          ) : (
            <div className="animate-spin">Loading ...</div>
          )}
        </Disclosure.Panel>
      </Transition>
    </Disclosure>
  );
}
