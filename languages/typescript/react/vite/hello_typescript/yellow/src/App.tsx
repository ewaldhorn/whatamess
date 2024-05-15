import { FC } from 'react'
import './App.css'

interface AppProperties {
  title: string;
}

const App: FC<AppProperties> = ({ title }) => {
  return <div>
    <h1>{title}</h1>
    And this is my app.
  </div>;
}

export default App
