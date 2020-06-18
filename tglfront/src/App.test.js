import React from 'react';
// import { render } from '@testing-library/react';
import { shallow, mount } from 'enzyme';
import App from './App';
import renderer from 'react-test-renderer';

test('checking jest/enzyme are working', () => {
  const num = 5;
  expect(2 + 3).toEqual(num);
})

describe('App', () => {
  let wrapper = shallow(<App />);

  it('shallow renders App', () => {
    wrapper;
  });

  it('Should render TaskComp', () => {
    expect(wrapper.find('TaskComp').length).toEqual(1);
  });
  
  it('Shouldn\'t render LogInComp', () => {
    expect(wrapper.find('LogInComp').length).toEqual(0);
  })

  // write test cases of what component should
  // render depending on the state values.
});


// snapshot testing

it('Renders App with TaskComp', () => {
  const AppComponent = renderer.create(<App />).toJSON();
  expect(AppComponent).toMatchSnapshot();
});






