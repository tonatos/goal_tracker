import React from 'react';
import { useEffect } from 'react';
import { RecoilRoot } from 'recoil'
import { MainContainer } from './containers/MainContainer';

import WebFont from 'webfontloader';

function App() {
  useEffect(() => {
    WebFont.load({
      google: {
        families: ['Kanit:500,700']
      }
    });
   }, []);

  return (
    <RecoilRoot>
      <MainContainer />
    </RecoilRoot>
  );
}

export default App;
