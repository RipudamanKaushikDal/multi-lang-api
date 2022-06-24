import './App.css';
import AppContextProvider from './context/context-provider';
import LandingPage from './landing-page/landing-page';

const App = () => {
  
  return (
    <div className="App">
      <AppContextProvider>
        <LandingPage />
      </AppContextProvider>  
    </div>
  );
}

export default App;
