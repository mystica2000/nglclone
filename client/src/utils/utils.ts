import store from "../store/store"

const { responses, setResponses } = store;

/** Sort by new based on the ID */
export const sortByNew = () => {
  const updatedResponses = [...responses].sort((a, b) => a.id - b.id)
  setResponses(updatedResponses)
}

/** Sort by not replied by whether the response.done is 1 or 0 */
export const sortByNotReplied = () => {
  const updatedResponse = [...responses].sort((a, b) => a.done - b.done);
  setResponses(updatedResponse)
}


export const sortHelper = (val: string) => {
  switch (val) {
    case "new": {
      sortByNew();
      break;
    }
    case "replied": {
      sortByNotReplied()
      break;
    }
    default: {
      break;
    }
  }
}