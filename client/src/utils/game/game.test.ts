import { expect, test } from 'vitest'

import {
  createWhiteChecker,
  createBlackChecker,
  createDefaultBoard,
  getTrianglePath,
  getTriangleColor,
  getCheckerX,
  BOARD,
  getCheckerY,
} from './game'

test('create white checker', () => {
  const position = 30
  const stackIndex = 12
  const checker = createWhiteChecker(position, stackIndex)
  expect(checker).toEqual({
    color: 'white',
    position,
    stackIndex,
  })
})

test('create black checker', () => {
  const position = 5
  const stackIndex = 2
  const checker = createBlackChecker(position, stackIndex)
  expect(checker).toEqual({
    color: 'black',
    position,
    stackIndex,
  })
})

test('create default board', () => {
  const board = createDefaultBoard()

  expect(board).toHaveLength(30)

  // Test initial checker positions
  expect(
    board.filter(c => c.color === 'white' && c.position === 0),
  ).toHaveLength(2)
  expect(
    board.filter(c => c.color === 'white' && c.position === 16),
  ).toHaveLength(3)
  expect(
    board.filter(c => c.color === 'white' && c.position === 18),
  ).toHaveLength(5)
  expect(
    board.filter(c => c.color === 'white' && c.position === 11),
  ).toHaveLength(5)

  expect(
    board.filter(c => c.color === 'black' && c.position === 23),
  ).toHaveLength(2)
  expect(
    board.filter(c => c.color === 'black' && c.position === 7),
  ).toHaveLength(3)
  expect(
    board.filter(c => c.color === 'black' && c.position === 5),
  ).toHaveLength(5)
  expect(
    board.filter(c => c.color === 'black' && c.position === 12),
  ).toHaveLength(5)
})

test('get triangle path', () => {
  // Upper triangle (position < 12)
  const upperPath = getTrianglePath(0)
  expect(upperPath).toContain('M ')
  expect(upperPath).toContain('L ')
  expect(upperPath).toContain('Z')

  // Lower triangle (position >= 12)
  const lowerPath = getTrianglePath(12)
  expect(lowerPath).toContain('M ')
  expect(lowerPath).toContain('L ')
  expect(lowerPath).toContain('Z')
})

test('get checker X position', () => {
  // Upper right
  expect(getCheckerX(0)).toBe(
    BOARD.padding +
      11 * BOARD.triangleWidth +
      BOARD.triangleWidth / 2 +
      BOARD.centerBarWidth * 2,
  )

  // Upper left
  expect(getCheckerX(11)).toBe(BOARD.padding + BOARD.triangleWidth / 2)

  // Lower left
  expect(getCheckerX(12)).toBe(BOARD.padding + BOARD.triangleWidth / 2)

  // Lower right
  expect(getCheckerX(23)).toBe(
    BOARD.padding +
      11 * BOARD.triangleWidth +
      BOARD.triangleWidth / 2 +
      BOARD.centerBarWidth * 2,
  )
})

test('get checker Y position', () => {
  // Upper triangle
  expect(getCheckerY(0, 0)).toBe(BOARD.padding + BOARD.checkerRadius)
  expect(getCheckerY(0, 1)).toBe(
    BOARD.padding + BOARD.checkerRadius + BOARD.checkerRadius * 1.8,
  )

  // Lower triangle
  expect(getCheckerY(12, 0)).toBe(
    BOARD.height - BOARD.padding - BOARD.checkerRadius,
  )
  expect(getCheckerY(12, 1)).toBe(
    BOARD.height -
      BOARD.padding -
      BOARD.checkerRadius -
      BOARD.checkerRadius * 1.8,
  )
})
