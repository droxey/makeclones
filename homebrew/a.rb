require 'rbconfig'
class A < Formula
  desc ""
  homepage "https://github.com/–desc==Clone/a"
  version "1"

  if Hardware::CPU.is_64_bit?
    case RbConfig::CONFIG['host_os']
    when /mswin|msys|mingw|cygwin|bccwin|wince|emc/
      :windows
    when /darwin|mac os/
      url "https://github.com/–desc==Clone/a/releases/download/v1/makeclones_1_darwin_amd64.zip"
      sha256 "aa0cce3bb45035698ec40a888a4e8ce5fb375685e083f8a19ca17f23e67c4a6d"
    when /linux/
      url "https://github.com/–desc==Clone/a/releases/download/v1/makeclones_1_linux_amd64.tar.gz"
      sha256 "76c5754a130248f5a134acda866b4629f5cff433e1ac7bc932233754659f184a"
    when /solaris|bsd/
      :unix
    else
      :unknown
    end
  else
    case RbConfig::CONFIG['host_os']
    when /mswin|msys|mingw|cygwin|bccwin|wince|emc/
      :windows
    when /darwin|mac os/
      url "https://github.com/–desc==Clone/a/releases/download/v1/makeclones_1_darwin_386.zip"
      sha256 "0083de3457831110dcf8eeb5b21f1c5ad4431cbb3a28d6041dc17ec49cef0d15"
    when /linux/
      url "https://github.com/–desc==Clone/a/releases/download/v1/makeclones_1_linux_386.tar.gz"
      sha256 "1c6a3c8759d8e6a58a4f485df3aa1b163b6fa53cc896e62680eb4331640381e3"
    when /solaris|bsd/
      :unix
    else
      :unknown
    end
  end

  def install
    bin.install "a"
  end

  test do
    system "a"
  end

end
