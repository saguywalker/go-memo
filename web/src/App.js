import React, { useState, useEffect } from "react";
import NoteBox from "./components/NoteBox";
import NoteInput from "./components/NoteInput";

import axios from 'axios';

const endpoint = "http://localhost:3000";

function App() {
  const [notes, setNotes] = useState([]);

  useEffect(() => {
    const fetchData = async () => {
      const result = await axios({
        method: 'get',
        url: endpoint+"/api/notes",
      });
      if (result.data == null) return;
      setNotes(result.data);
    };

    fetchData();
  }, []);

  function addNote({ title, noteDetail }) {
    axios.post(
      endpoint+"/api/notes", 
      { id: 0, title: title, detail: noteDetail, lastEdit: "" },
      {
        headers: {
          "Content-Type": "application/json"
        },
        responseType: 'json'
      }
    ).then( (res) => {
      const newNotes = [...notes, res.data];
      setNotes(newNotes);
    });
    // const result = await axios(endpoint+"/api/notes")
    // const newNotes = [...notes, { title, noteDetail }];
    // setNotes(newNotes);
  };

  function setEditNote(editValue, index) {
    const payload = {id: index, title: editValue.title, detail: editValue.noteDetail, lastEdit: +new Date}
    axios.put(
      endpoint+"/api/notes", 
      {editValue},
      {
        headers: {
          "Content-Type": "application/json"
        },
        responseType: 'json'
      }
    ).then( (res) => {
      const newNotes = [...notes, res.data];
      setNotes(newNotes);
    });
    // const newNotes = [...notes];
    // newNotes[index] = editValue;
    // setNotes(newNotes);
  }

  return (
    <div className="App">
      <NoteInput 
        addNote={addNote}
      />
      <h2 style={{color: 'white'}}>My Notes</h2>
      <div className="Note-list">
        {
          notes.length !== 0 ? 
            notes.map((note, index) => (
              <NoteBox
                key={index}
                index={index}
                note={note}
                setEditValueCallbackParent={(edit) => setEditNote(edit, index)}
              />
            )) :
            <h4>There is no note.</h4>  
        }
      </div>
    </div>
  );
}

export default App;