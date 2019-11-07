import React, { useState, useEffect } from "react";
import NoteBox from "./components/NoteBox";
import NoteInput from "./components/NoteInput";

import axios from 'axios';

const endpoint = "http://localhost:3000";

function App() {
  const [notes, setNotes] = useState([]);

  // Fetch data from db first time on load
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

  function addNote(editValue) {
    const payload =  {
      title: editValue.title,
      detail: editValue.detail
    };
    axios.post(
      `${endpoint}/api/notes`,
      payload,
      {
        headers: {
          "Content-Type": "application/json"
        },
        responseType: 'json'
      }
    ).then((res) => {
      const newNotes = [...notes, res.data];
      setNotes(newNotes);
    });
  };

  function setEditNoteMain(editValue) {
    const payload = {
      title: editValue.title,
      detail: editValue.detail,
      lastEdit:  new Date()
    };
    axios.put(
      `${endpoint}/api/notes/${editValue.id}`, 
      payload,
      {
        headers: {
          "Content-Type": "application/json"
        },
        responseType: 'json'
      }
    ).then((res) => {
      setNotes(notes.map((note) => {
        if (note.id === res.data.id) return res.data;
        return note;
      }));
    });
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
                setEditValueCallbackParent={(edit) => setEditNoteMain(edit)}
              />
            )) :
            <h4>There is no note.</h4>  
        }
      </div>
    </div>
  );
}

export default App;