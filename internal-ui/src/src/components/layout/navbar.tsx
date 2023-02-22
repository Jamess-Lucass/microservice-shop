import {
  Flex,
  IconButton,
  useColorModeValue,
  Text,
  HStack,
  Menu,
  MenuButton,
  Avatar,
  VStack,
  Box,
  MenuList,
  MenuItem,
  MenuDivider,
  useColorMode,
  Link,
  Icon,
  Image,
} from "@chakra-ui/react";
import { FiChevronDown, FiLogOut, FiMenu, FiMoon } from "react-icons/fi";

type NavbarProps = {
  onMobileSidebarToggle: () => void;
};

export function Navbar({ onMobileSidebarToggle }: NavbarProps) {
  const { toggleColorMode } = useColorMode();

  return (
    <Flex
      pr={4}
      pl={2}
      height={16}
      alignItems="center"
      borderBottomWidth="1px"
      bg={useColorModeValue("white", "gray.900")}
      borderBottomColor={useColorModeValue("gray.200", "gray.700")}
      justifyContent="space-between"
    >
      <Flex alignItems="center" gap={1}>
        <IconButton
          onClick={onMobileSidebarToggle}
          display={{ base: "inline-flex", sm: "none" }}
          variant="ghost"
          aria-label="open menu"
          isRound
          icon={<FiMenu />}
        />

        <Link href="/" display={{ base: "none", sm: "block" }}>
          <Image
            src={useColorModeValue("/logo-dark.svg", "/logo-light.svg")}
            h={8}
            alt="Logo"
          />
        </Link>
      </Flex>

      <HStack spacing={{ base: "2", md: "6" }}>
        <Flex alignItems="center">
          <Menu isLazy>
            <MenuButton
              py={2}
              transition="all 0.3s"
              _focus={{ boxShadow: "none" }}
            >
              <HStack>
                <Avatar size="sm" />
                <VStack
                  display={{ base: "none", md: "flex" }}
                  alignItems="flex-start"
                  spacing="1px"
                >
                  <Text fontSize="sm" textTransform="capitalize"></Text>
                  <Text
                    fontSize="xs"
                    color="gray.600"
                    textTransform="capitalize"
                  ></Text>
                </VStack>
                <Box display={{ base: "none", md: "flex" }}>
                  <FiChevronDown />
                </Box>
              </HStack>
            </MenuButton>
            <MenuList fontSize="sm">
              <MenuItem onClick={toggleColorMode}>
                <Icon as={FiMoon} mx={2} />
                Toggle theme
              </MenuItem>

              <MenuDivider />

              <MenuItem>
                <Icon as={FiLogOut} mx={2} />
                Sign out
              </MenuItem>
            </MenuList>
          </Menu>
        </Flex>
      </HStack>
    </Flex>
  );
}
