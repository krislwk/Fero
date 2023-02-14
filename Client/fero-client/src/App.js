import './App.css';
import AddButton from "./components/AddButton";
import Header from "./components/Header";
import SaveButton from "./components/SaveButton";

function App() {
  return (
    <div className="App">
        <Header prop = "Kris"/>
        <SaveButton/>
        <AddButton/>

    </div>
  );
}

export default App;
