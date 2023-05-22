import { Disclosure, Transition } from '@headlessui/react'
import { useState } from 'react';

export default function MyDisclosure() {
    const [isFetched, setIsFetched] = useState(false);
    let versionText = "";
    const fetchedData = () => {
        window.fetch("https://raw.githubusercontent.com/tumGER/Kosmonaut/gh-pages/assets/version.txt")
        .then((response) => {
            if (!response.ok) {
                return "Error\nGithub might be having problems"
            }
            return response.text();
        })
        .then((text) => {
            setIsFetched(true);
            versionText = text;
        })
    };
  return (
    <Disclosure>
      <Disclosure.Button className="font-light text-lg" onClick={fetchedData}>
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
        {
            isFetched ?
            (<div className='whitespace-pre-line'>{versionText == "" ? "Error!\nIt seems like the version data wasn't generated" : versionText}</div>) :
            (<div className='animate-pulse'>Fetching Data ...</div>)
        }
      </Disclosure.Panel>
      </Transition>
    </Disclosure>
  )
}