import React from 'react';

import Header from './components/Header'
import TextField from '@material-ui/core/TextField';
import IconButton from '@material-ui/core/IconButton';
import Button from '@material-ui/core/Button';
import DeleteIcon from '@material-ui/icons/Delete';

function App() {
  return (
    <>
      <Header title="" />
      <TextField id="outlined-basic" label="Company Name" variant="outlined" />
      <TextField id="outlined-basic" label="Ticker" variant="outlined" />
      <Button variant="contained">Add</Button>
      <IconButton aria-label="delete">
        <DeleteIcon />
      </IconButton>
    </>
  );
}

export default App;
