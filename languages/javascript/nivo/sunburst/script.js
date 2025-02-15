// script.js
const data = [
  {
    id: 'root',
    children: [
      {
        id: 'herbs',
        children: [
          { id: 'basil', value: 10 },
          { id: 'thyme', value: 15 },
          { id: 'rosemary', value: 20 },
        ],
      },
      {
        id: 'spices',
        children: [
          { id: 'cinnamon', value: 12 },
          { id: 'nutmeg', value: 18 },
          { id: 'ginger', value: 22 },
        ],
      },
    ],
  },
];

const chart = document.getElementById('chart');
const sunburst = new Nivo.Sunburst(chart, {
  data,
  margin: { top: 50, right: 50, bottom: 50, left: 50 },
  id: 'id',
  value: 'value',
  cornerRadius: 2,
  borderColor: { from: 'color', modifiers: [] },
  colors: { scheme: 'nivo' },
  childColor: { from: 'color', modifiers: [] },
  animate: true,
  motionConfig: 'stiff',
});
