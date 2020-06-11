import React from 'react';
// import { render } from '@testing-library/react';
import { shallow } from 'enzyme';
import App from './App';

test('checking jest/enzyme are working', () => {
  const num = 5;
  expect(2 + 3).toEqual(num);
})

describe('App', () => {
  let wrapper = shallow(<App />);

  it('shallow renders App', () => {
    wrapper;
  });

  it('Should have 1 div', () => {
    expect(wrapper.find('div').length).toEqual(1);
  });

  // write test cases of what component should
  // render depending on the state values.
});



