import React, { useState, useEffect } from "react";
import NoteBox from "./components/NoteBox";
import NoteInput from "./components/NoteInput";
import {FetchData, Store, Update, FetchByID} from "./services";

function App() {
  const [notes, setNotes] = useState([]);

  useEffect(() => {
    FetchData().then((res) => {
      setNotes(res.data)
    });
  }, []);
/*
  const addNote = (note) => {
    Store(note).then((res) => {
      setNotes([...notes, res.data])
    });
  };
*/

  const addNote = async (note) => {
    const resp = await Store(note);
    setNotes([...notes, resp.data]);
  };

  const setEditNoteMain = async (editValue) => {
    const resp = await Update(editValue);
    const editValueIndex = notes.indexOf(notes.find(note => note.id === editValue.id));
    notes[editValueIndex] = resp.data;
    setNotes([...notes]);
  };

  const fetchNoteById = async (noteId) => {
    const resp = await FetchByID(noteId);
    const noteIndex = notes.indexOf(notes.find(note => note.id === noteId));
    notes[noteIndex] = resp.data;
    setNotes([...notes]);
    return resp.data;
  };

  return (
    <div className="App">
      <NoteInput 
        addNote={addNote}
      />
      <h2 style={{color: 'white'}}>Go-Memo</h2>
      <div className="Note-list">
        {
          notes.length !== 0 ? 
            notes.map((note) => (
              <NoteBox
                key={note.id}
                note={note}
                fetchNoteByIdCallbackToParent={(id) => fetchNoteById(id)}
                setEditValueCallbackToParent={(edit) => setEditNoteMain(edit)}
              />
            )) :
            <h4>There is no note.</h4>
           
        }
      </div>
    </div>
  );
}

export default App;