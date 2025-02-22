/**
 * Creating a sidebar enables you to:
 - create an ordered group of docs
 - render a sidebar for each doc of that group
 - provide next/previous navigation

 The sidebars can be generated from the filesystem, or explicitly defined here.

 Create as many sidebars as you want.
 */

// @ts-check

/** @type {import('@docusaurus/plugin-content-docs').SidebarsConfig} */
const sidebars = {
  yarasec: [
    {
      type: 'html',
      value: 'Khulnasoft YaraSec',
      className: 'sidebar-title',
    },    

    "yarasec/index",
    "yarasec/quickstart",

    {
      type: 'category',
      label: 'Using YaraSec',
      items: [
        'yarasec/using/build',
        'yarasec/using/scan',
      ]
    },

    {
      type: 'category',
      label: 'Configuration',
      items: [
        'yarasec/configure/cli',
        'yarasec/configure/output',
        'yarasec/configure/rules',
      ]
    },

  ],
};

module.exports = sidebars;
