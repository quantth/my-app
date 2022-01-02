const reducer = (state = {}, action) => {
    switch (action.type) {
      case 'PARSE_POST_TITLE':
        return { ...state, loading: true };
      default:
        return state;
    }
  };
  
  export default reducer;