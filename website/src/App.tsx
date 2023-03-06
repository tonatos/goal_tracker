import React from 'react';
import { RecoilRoot } from 'recoil'
import { MainContainer } from './containers/MainContainer';

function App() {
  return (
    <RecoilRoot>
      <MainContainer />
    </RecoilRoot>
  );
}

export default App;
